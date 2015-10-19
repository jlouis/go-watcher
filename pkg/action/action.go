package action

import (
	"fmt"
	"github.com/shopgun/"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/giulioungaretti/go-xz"
	s3 "github.com/giulioungaretti/s3Manager"
)

func makeRelativePath(path string, base string) string {
	pathSplit := strings.Split(path, "/")
	flatPath := pathSplit[len(pathSplit)-1 : len(pathSplit)]
	relativepathSplit := strings.Split(strings.Join(flatPath, ""), "_")
	relativepath := strings.Join(relativepathSplit, "/")
	newmsgpackPath := strings.Join([]string{base, relativepath}, "/")
	return newmsgpackPath
}

func zeropad(path string) string {
	pathSplit := strings.Split(path, "/")
	scope := pathSplit[len(pathSplit)-6]
	action := pathSplit[len(pathSplit)-5]
	year := pathSplit[len(pathSplit)-4]
	month := pathSplit[len(pathSplit)-3]
	day := pathSplit[len(pathSplit)-2]
	if len(day) < 2 {
		day = fmt.Sprintf("0%v", day)
	}
	if len(month) < 2 {
		month = fmt.Sprintf("0%v", month)
	}
	filename := pathSplit[len(pathSplit)-1]
	pieces := make([]string, 0)
	pieces = append(pieces, scope)
	pieces = append(pieces, action)
	pieces = append(pieces, year)
	pieces = append(pieces, month)
	pieces = append(pieces, day)
	pieces = append(pieces, filename)
	path = strings.Join(pieces, "/")
	return path
}

type XzToS3 struct {
	Conn *s3.Connection
}

func (a XzToS3) Connect() error {
	c := s3.Connection{}
	err := c.Connect()
	if err != nil {
		return err
	}
	a.Conn = &c
	return nil
}

func (a XzToS3) Do(event string) error {
	//
	bucket, err := s3.Getbucket("eta-events-msgpack", a.Conn)
	msg := serialize.Msgpack{}
	// now compress  and check integrity
	err = xz.DeflateCheck(event, "sha256")
	if err != nil {
		log.Errorf("error deflating %v, %v", event, err)
	}
	// now file is on disk
	xzcompressedpath := fmt.Sprintf("%v.xz", event)
	bytes, err := ioutil.ReadFile(xzcompressedpath)
	if err != nil {
		log.Errorf("", event, err)
		fmt.Printf(err.Error())
	} else {
		msg.Bytes = bytes
		msg.Path = zeropad(makeRelativePath(xzcompressedpath, ""))
		msg.Checksum.Md5 = xz.Base64md5(bytes)
	}
	err = s3.Put(bucket, msg)
	if err != nil {
		log.Warnf("error uploading to s3 :%v", err)
	}
	fmt.Printf("sent:%v \n", xzcompressedpath)
	// remove file
	err = os.Remove(xzcompressedpath)
	if err != nil {
		log.Errorf("%v, %v", event, err)
	}
	fmt.Printf("Removed:%v \n", xzcompressedpath)
	return nil
}

func S3Mock(event string) error {
	// now file is on disk
	xzcompressedpath := fmt.Sprintf("%v.xz", event)
	fmt.Printf("fake sent:%v \n", zeropad(makeRelativePath(xzcompressedpath, "")))
	return nil
}
