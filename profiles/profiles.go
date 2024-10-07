package profiles

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

type Profile struct {
	FileLocation 	string
	Using			bool
}

