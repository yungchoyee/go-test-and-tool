package password

import (
	"math/rand"
	"time"
)

type PassType int //密码组合类型

const (
	PassTypeNumAndLetter PassType = 1
	PassTypeAllNum PassType = 2
	PassTypeAllLetter PassType = 3
	PassTypeAddOperator PassType = 4
	PassTypeAddSpecialChar PassType = 5 //+-*/!@#$%^&*()_=
	PassTypeAddDefine  PassType = 6
)
//当前密码对象
type pass struct {
	Length int
	PT PassType
	DefineChar []string
}

func NewPass(length int,passType int,defineChar... string)(obj *pass){
	obj = new(pass)
	obj.Length = length
	obj.PT = PassType(passType)
	obj.DefineChar = defineChar
	return
}
//获取组成密码的字符
func (p *pass)createPassChars()(chars []string){
	switch p.PT {
	case PassTypeNumAndLetter:
		chars = getLetterAndNum()
	case PassTypeAllNum:
		chars = getAllNum()
	case PassTypeAllLetter:
		chars = getAllLetter()
	case PassTypeAddOperator:
		chars = getAddOperator()
	case PassTypeAddSpecialChar:
		chars = getAddSpecial()
	case PassTypeAddDefine:
		chars = p.getDefineChars()
	}
	return
}

//all number
func getAllNum()[]string{
	return []string{"0","1","2","3","4","5","6","7","8","9"}
}

//all letter
func getAllLetter()[]string{
	return []string{
		"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z",
		"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
	}
}
//add operator
func getAddOperator()[]string{
	return  append(getLetterAndNum(),getOperator()...)
}
func getOperator()[]string{
	return []string{"+","-","*","/"}
}
//add special char
func getAddSpecial()[]string{
	return append(getLetterAndNum(),getSpecialChars()...)
}
func getSpecialChars()[]string{
	return []string{"!","@","#","$","%","^","&","(",")"}
}

//add define chars
func (p *pass)getDefineChars()(chars []string){
	chars = getLetterAndNum()
	chars = append(chars,p.DefineChar...)
	return
}

//number + letter
func getLetterAndNum()[]string{
	return append(getAllLetter(),getAllNum()...)
}
//create password
func (p *pass)CreatePassword () (password string){
	chars := p.createPassChars()
	rand.Seed(time.Now().Unix())
	for len(password) < p.Length - 1{
		password += chars[rand.Intn(len(chars))]
	}
	if p.PT != PassTypeAllNum {
		password = getAllLetter()[rand.Intn(1)]
	}
	return
}
