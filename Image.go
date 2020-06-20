package main 
import(
	"fmt"
	"os"
	"bufio"
	"log"
	"image"
	_ "image/jpeg"
	_ "image/draw"
	_ "image/color"

)
func main(){

	fmt.Println("Merci d'écrire le nom du fichier contenant l'image. \nPar exemple : image.jpeg")
	reader:=bufio.NewReader(os.Stdin)
	text,_,_ :=reader.ReadLine()
	txt := string(text)
	fmt.Println(txt)
	decomp(txt)
	//construct()
	//
}
func decomp(filename string) ([]uint32,[]uint32,[]uint32,[]uint32){

	reader, err:=os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer reader.Close()
	im,_,err := image.Decode(reader)
	b := im.Bounds()
	fmt.Println(im.At(10,10).RGBA())
	R := []uint32{}
	G := []uint32{}
	B := []uint32{}
	A := []uint32{}
	for j := b.Min.Y; j< b.Max.Y;j++{
		for i := b.Min.X; i<b.Max.X;i++{
			//fmt.Println(im.At(i,j).RGBA())
			r,g,b,a := im.At(i,j).RGBA()
			//temp := color.RGBA(r,g,b,a)
			R = append(R,r)
			G = append(G,g)
			B = append(B,b)
			A = append(A,a)
			fmt.Println(r,g,b,a)
			//a = 2
			//im.At(i,j).RGBA() = r,g,b,a
		}
	}
	fmt.Println("HERE WE GO",R)
	//out, err := os.Create("output.jpeg")
	//err = jpeg.Encode(out,im,80)
	fmt.Println("Image générée ")
	return R,G,B,A
}
func filterLumi(R,G,B,A []uint32)([]uint32,[]uint32,[]uint32,[]uint32){
	chR := make(chan uint32)
	chG := make(chan uint32)
	chB := make(chan uint32)
	go goLumi(R,100,chR)
	go goLumi(G,100,chG)
	go goLumi(B,100,chB)
	return retR,retG,retB,A

}
func goLumi(R []uint32, l uint32, ch chan uint32,){
	for i:=0;i<len(R);i++{
		ch <- R[i]+l
	}
}
