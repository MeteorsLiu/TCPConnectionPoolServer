package main
import (
"log"
"net"

"encoding/binary"
)
func serve() {
	l, err := net.Listen("tcp", "127.0.0.1:9998")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go func() {
			defer c.Close()
			b := make([]byte, 2)
			//res, _ := io.ReadAll(c)
			//b := make([]byte, 2
			for {
				if n, err := c.Read(b[:2]); n == 0 || err != nil {
					log.Printf("%d Exit", binary.LittleEndian.Uint16(b))
					return
				}
				log.Println(binary.LittleEndian.Uint16(b))
			}	
			
		}()
	}
}

func main(){
serve()
}
