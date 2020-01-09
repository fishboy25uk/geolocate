package geolocate

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/oschwald/maxminddb-golang"
)

var (
	dbCity *maxminddb.Reader
	dbASN  *maxminddb.Reader

	dbASNFileName  = "GeoLite2-ASN.mmdb"
	dbCityFileName = "GeoLite2-City.mmdb"

	dbASNFileNameCompressed  = "GeoLite2-ASN.tar.gz"
	dbCityFileNameCompressed = "GeoLite2-City.tar.gz"
	//dbASNMD5Path             = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN.tar.gz.md5"
	//dbCityMD5Path            = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5"
	//dbCityDBPath             = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"
	//dbASNDBPath              = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN.tar.gz"

	dbASNDBPath   = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&license_key=%s&suffix=tar.gz"
	dbASNMD5Path  = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&license_key=%s&suffix=tar.gz.md5"
	dbCityDBPath  = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=%s&suffix=tar.gz"
	dbCityMD5Path = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=%s&suffix=tar.gz.md5"
)

//OpenDBs opens the MaxMind City and ASN databases
func OpenDBs() {

	//Open GeoIP City DB
	dbCityTemp, err := openDB(dbCityFileNameCompressed, dbCityFileName)
	if err != nil {
		log.Println("ERROR: Could not load geoip dbCity", err)
		os.Exit(1)
	}
	dbCity = dbCityTemp
	//defer dbCity.Close()

	//Open GeoIP ASN DB
	dbASNTemp, err := openDB(dbASNFileNameCompressed, dbASNFileName)
	if err != nil {
		log.Println("ERROR: Could not load geoip dbASN", err)
		os.Exit(1)
	}
	dbASN = dbASNTemp
	//defer dbASN.Close()

}

func openDB(filenameCompressed string, filename string) (*maxminddb.Reader, error) {

	file, err := os.Open(filenameCompressed)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	tr := tar.NewReader(gzReader)
	for {

		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}

		if filepath.Base(hdr.Name) == filename {

			dbBytes, err := ioutil.ReadAll(tr)
			if err != nil {
				return nil, err
			}

			dbMM, err := maxminddb.FromBytes(dbBytes)
			if err != nil {
				return nil, err
			}

			return dbMM, nil
		}

	}

	return nil, errors.New("DB file not found in archive")

}
