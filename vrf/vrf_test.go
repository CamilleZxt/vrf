package vrf

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"io"
	"testing"
)

func TestVrf(t *testing.T) {
	for i := 0; i < 10; i++ {
		sk, err := ecdsa.GenerateKey(curve, rand.Reader)
		if nil != err {
			t.Fatal("GenerateKey fail", err)
		}
		sk2, err := ecdsa.GenerateKey(curve, rand.Reader)
		if nil != err {
			t.Fatal("GenerateKey fail", err)
		}
		data := make([]byte, 32)
		io.ReadFull(rand.Reader, data)
		t.Log("data:", hex.EncodeToString(data))
		pi, err := Prove(sk, data)
		if nil != err {
			t.Fatal("Generate vrf proof failed", err)
		}
		// 用对应的PublicKey和data验证
		ok, err := Verify(&sk.PublicKey, pi, data)
		t.Log("ok:", ok)      //true
		t.Log("ok err:", err) //nil
		if nil != err || !ok {
			t.Fatal("verification failed", err)
		}
		// 用不对应的PublicKey验证
		ok2, err := Verify(&sk2.PublicKey, pi, data)
		t.Log("ok2:", ok2)     //false
		t.Log("ok2 err:", err) //nil
		if nil != err || ok2 {
			t.Fatal("verification failed", err)
		}
		// 用不对应的data验证
		data = append(data, []byte("message")...)
		ok3, err := Verify(&sk.PublicKey, pi, data)
		t.Log("ok3:", ok3)     //false
		t.Log("ok3 err:", err) //nil
		if nil != err || ok3 {
			t.Fatal("verification failed", err)
		}
	}
}

func TestVrfElection(t *testing.T) {
	var nonce []byte
	VrfElection(nonce)
}
