package go_say_hello

// go modules
// gunakan v untuk membuat tag di terminal git bash. contoh : v1.0.0
// cara membuat tag : git tag v1.0.0
// kemudian lakukan push : git push origin v1.0.0

// updgrade module
// bisa dilakukan dengan melakukan perubahan terlebih dahulu
// kemudian dilakukan commit, dan melakukan push tag yang baru

// major upgrade
// biasanya dilakukan ketika terjadi upgrade modules yang besar
// hindari tindakan upgrade major dengan nama modules yang sama
// sehingga nanti cukup di ubah saja, semisal mau upgrade major ke versi diatasnya yaitu v2.0.0
// maka nama modules di project cukup diubah seperti ini :
// github.com/Akuutauf-Portofolio/go-say-hello/v2
// supaya nanti orang yang masih menggunakan versi yang sebelum upgrade major masih bisa mengunakannya
// tujuannya agar tidak terjadi backward compatible: yang berpengaruh ke kode lainnya

func SayHello(name string) string {
	return "Hello " + name
}