package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"github.com/lxn/walk"
	"log"
	"os"
	"path/filepath"
)

func genFile() {
	var t walk.Form

	if url.Text() == "" {
		walk.MsgBox(t, "提示", "URL不能为空！！", walk.MsgBoxIconWarning)
		return
	}

	if user.Text() == "" {
		walk.MsgBox(t, "提示", "用户名不能为空!!", walk.MsgBoxIconWarning)
		return
	}

	if password.Text() == "" {
		walk.MsgBox(t, "提示", "密码不能为空!!", walk.MsgBoxIconWarning)
		return
	}

	if path.Text() == "" {
		walk.MsgBox(t, "提示", "保存路径不能为空!!", walk.MsgBoxIconWarning)
		return
	}

	//JDBC文件内容
	jdbc := "jdbc.url=" + url.Text() + "\n" + "jdbc.username=" + user.Text() + "\n" + "jdbc.password=" + password.Text() + "\n"
	log.Println(jdbc)
	err := encryptFile([]byte(jdbc))
	if err != nil {
		walk.MsgBox(t, "提示", "生成文件失败", walk.MsgBoxIconError)
		return
	}

	walk.MsgBox(t, "提示", "生成成功!!", walk.MsgBoxIconInformation)
	return
}

func encryptFile(jdbcstr []byte) error {
	var t walk.Form

	//java key  {-28, -72, -83, -27, -101, -67, -27, -92, -89, -28, -72, -128, -25, -69, -97, 33}
	//int8 to uint8 = 256+int8
	key := []byte{228, 184, 173, 229, 155, 189, 229, 164, 167, 228, 184, 128, 231, 187, 159, 33}

	//建立JDBC文件
	fp, err := os.OpenFile(filepath.Join(path.Text(), DEFAULT_FILE_NAME), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		walk.MsgBox(t, "提示", "建立文件失败", walk.MsgBoxIconError)
		return err
	}
	defer fp.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return err
	}
	blockSize := block.BlockSize()

	oridata := PKCS5Padding(jdbcstr, blockSize)

	dst := make([]byte, len(oridata))

	for i, count := 0, len(oridata)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(dst[begin:end], oridata[begin:end])
	}

	log.Println("encrypt:", hex.EncodeToString(dst))
	jdbc := base64.StdEncoding.EncodeToString(dst)

	fp.WriteString(jdbc)

	return nil
}

func decrypt(jdbcstr string) (plain string, err error) {
	//java key  {-28, -72, -83, -27, -101, -67, -27, -92, -89, -28, -72, -128, -25, -69, -97, 33}
	//int8 to uint8 = 256+int8
	key := []byte{228, 184, 173, 229, 155, 189, 229, 164, 167, 228, 184, 128, 231, 187, 159, 33}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	blockSize := block.BlockSize()
	jdbc, err := base64.StdEncoding.DecodeString(jdbcstr)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dst := make([]byte, len(jdbc))

	for i, count := 0, len(jdbc)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(dst[begin:end], jdbc[begin:end])
	}

	res := PKCS5Unpadding(dst)
	plain = string(res)

	return plain, nil
}
