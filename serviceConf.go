package lib

import (
	viperLib "github.com/qamarian-lib/viper"
	"github.com/spf13/afero"
)

func ServiceConf () (dbUser, dbPass string, connTimeout, writeTimeout, readTimeout uint16, dbmsPubKey string, err error) {
	conf, errX := viperLib.NewFileViper (serviceConf_ConfFileName, "yaml")
	if errX != nil {
		err = err.New ("Unable to load conf file.", nil, nil, errX)
		return
	}

	// Processing conf data 'dbms_user_name'.  ..1.. {
	if conf.GetString ("dbms_user_name") == "" {
		err = err.New ("Conf data 'dbms_user_name': Data not set.", nil, nil)
		return
	}
	dbUser = conf.GetString ("dbms_user_name")
	// .. }

	// Processing conf data 'dbms_user_pass'.  ..1.. {
	if conf.GetString ("dbms_user_pass") == "" {
		err = err.New ("Conf data 'dbms_user_pass': Data not set.", nil, nil)
		return
	}
	dbPass = conf.GetString ("dbms_user_pass")
	// .. }

	// Processing conf data 'conn_timeout'.  ..1.. {
	if conf.GetString ("conn_timeout") == "" {
		err = err.New ("Conf data 'conn_timeout': Data not set.", nil, nil)
		return
	}
	timeoutF, errF := strconv.Atoi (conf.GetString ("conn_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'conn_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'conn_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'conn_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	connTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'wrte_timeout'.  ..1.. {
	if conf.GetString ("wrte_timeout") == "" {
		err = err.New ("Conf data 'wrte_timeout': Data not set.", nil, nil)
		return
	}
	timeoutF, errF := strconv.Atoi (conf.GetString ("wrte_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'wrte_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'wrte_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'wrte_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	writeTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'read_timeout'.  ..1.. {
	if conf.GetString ("read_timeout") == "" {
		err = err.New ("Conf data 'read_timeout': Data not set.", nil, nil)
		return
	}
	timeoutF, errF := strconv.Atoi (conf.GetString ("read_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'read_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'read_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'read_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	readTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'dbms_pub_key'.  ..1.. {	
	if conf.GetString ("dbms_pub_key") == "" {
		err = err.New ("Conf data 'dbms_pub_key': Data not set.", nil, nil)
		return
	}
	dbmsPubKey = conf.GetString ("dbms_pub_key")
	okJ, errJ := afero.Exists (afero.NewOsFs (), dbmsPubKey)
	if errJ != nil {
		err = err.New ("Conf data 'dbms_pub_key': Unable to confirm existence of file.", nil, nil, errJ)
		return
	}
	if okJ == false {
		err = err.New ("Conf data 'dbms_pub_key': File not found.", nil, nil)
		return
	}
	// .. }
}
var serviceConf_ConfFileName = "serviceConf.yml"
