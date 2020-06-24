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
	R,G,B,A := decomp(txt)
	filterLumi(R,G,B,A)
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
			//a = 2
			//im.At(i,j).RGBA() = r,g,b,a
		}
	}
	fmt.Println("Avant decomp",R[0])
	//out, err := os.Create("output.jpeg")
	//err = jpeg.Encode(out,im,80)
	return R,G,B,A
}
func filterLumi(R,G,B,A []uint32)([]uint32,[]uint32,[]uint32,[]uint32){
	chR := make(chan uint32)
	chG := make(chan uint32)
	chB := make(chan uint32)
	fmt.Println("ici")
	go goLumi(R,100,chR)
	retR :=make([]uint32,0,len(R))
	okR := true
	for okR==true{
		var temp uint32
		temp, ok :=<- chR
		retR = append(retR,temp)
		if ok == false {
			okR = false
		}
	}
	go goLumi(G,100,chG)
	retG :=make([]uint32,0,len(G))
        okG := true
        for okG==true{
                var temp uint32
                temp, ok :=<- chG
                retG = append(retG,temp)
                if ok == false {
                        okG = false
                }
        }
	go goLumi(B,100,chB)
	retB :=make([]uint32,0,len(B))
        okB := true
        for okB==true{
                var temp uint32
                temp, ok :=<- chB
                retB = append(retB,temp)
                if ok == false {
                        okB = false
                }
        }
	fmt.Println(retR[0])
	return retR,retG,retB,A

}
func goLumi(R []uint32, l uint32, ch chan uint32,){
	for i:=0;i<len(R);i++{
		ch <- R[i]+l
	}
	close(ch)
}
