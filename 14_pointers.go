package main

import "fmt"

// pointer, değerimizin hafızadaki (memory) adresini tutar.
// pointerlarla ilk kez karşılaşanlar, akıllarında şu şekilde tutabilirler.
// A kişisi size bir evi, arabayı, adresi sorabilir. Parmağınızla evi gösterebilirsiniz
// parmağınız bu durumda ev, araba ya da sorulan adres değil, onlardan birini işaret eden araçtır. Yani "pointer"dır.
// aynı şekilde pointerın pointerını da oluşturabilirsiniz. yani diğer parmağınızla işaret ettiğiniz parmağı işaret edebilirsiniz.
// (tam adresini hatırlayamadığım ama stackoverflow'da okuduğumu düşündüğüm bir metafor)
func main() {
	myCar := "Black Mercedes"
	// & ile adresi işaret ederken, * ile işaretçimizin (pointer) içeriğini okuruz
	pointToMyCar := &myCar
	fmt.Println("Address of my lovely car:", pointToMyCar)
	fmt.Println("Car itself:", *pointToMyCar)
	*pointToMyCar = "Red Porsche"
	fmt.Println("My new car:", myCar)
}
