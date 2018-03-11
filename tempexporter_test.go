package main

import (
	"testing"
)

var testdata = `Number of devices: 0


SENSORS:

PROTOCOL            	MODEL               	ID   	TEMP 	HUMIDITY	RAIN                	WIND                	LAST UPDATED        
fineoffset          	temperaturehumidity 	135  	24.2°	26%!	(MISSING)                    	                    	2018-01-25 18:16:30 
fineoffset          	temperature         	136  	5.7°	        	                    	                    	2018-01-25 18:16:35`

func TestParseString(t *testing.T) {
	parseddata := ParseString(testdata)
	t.Error(parseddata)
}
