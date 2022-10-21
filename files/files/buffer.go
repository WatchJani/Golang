package main

import (
	"bufio"
	"log"
	"os"
)

func Check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	file, err := os.OpenFile("d1/file.txt", os.O_WRONLY, 066)
	Check(err)

	defer file.Close()

	bufferWriter := bufio.NewWriter(file)

	bytesWritten, err := bufferWriter.Write([]byte{65, 66, 67}) //POSTOJI SANSA DA OVO NECU NIKAD IMATI :D

	log.Println(bytesWritten)

	bytesWritten, err = bufferWriter.WriteString("Buferred string\n") //OVO CU DA KORISTIM CESCE KOD ZAPISA U FILE
	log.Println(bytesWritten)

	unflushedBufferSize := bufferWriter.Buffered() //KOLIKO JE PROSTORA ZAUZETO U BAFERU
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	bytesAvaible := bufferWriter.Available() //KOLIKO JE PROSTORO OSTALO NA BAFERU
	log.Printf("Avaible buffer :%d", bytesAvaible)

	bufferWriter.Flush() //PISANJE PODATAKA IZ BAFERA U FILE
	bufferWriter.Reset(bufferWriter)

	unflushedBufferSize = bufferWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	bufferWriter = bufio.NewWriterSize(bufferWriter, 8000) //SETOVANJE NOVE VRIJEDNOSTI BAFERA (NAJMANJE JE 4KB(4096)BAJT)

	bytesAvaible = bufferWriter.Available()
	log.Println(bytesAvaible)

}
