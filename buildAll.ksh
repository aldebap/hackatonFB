#!	/usr/bin/ksh

export	PROJECT_PATH=$( cd $( dirname $0 ); pwd )
export	PROJECT=$( echo ${PROJECT_PATH} | sed -e 's/^.*\///g' )

echo ">>>>> Building \"${PROJECT}\" at ${PROJECT_PATH}"
echo

function buildGoLang {

	export  COMPONENT=$1
	export  GOPATH=${PROJECT_PATH}/${COMPONENT}
	export	GOOS=linux

	echo ">>> Building component \"${COMPONENT}\" at ${GOPATH}"
	echo

	#	get all required packages

	for SOURCE in $( find ${GOPATH}/src -name "*.go" | grep -v --regexp gopkg --regexp github )
	do
		for PACKAGE in $( grep --regexp=gopkg --regexp=github ${SOURCE} | sed -e 's/^[^"]*"/"/' | tr '"' ' ' )
		do
			echo "> Fetching required package ${PACKAGE}"

			go get ${PACKAGE}
		done
	done

	#	build the component

	echo "> Building source code"

	go build -o ${GOPATH}/bin/${COMPONENT} $( ls ${GOPATH}/src/*.go )

	echo
}

function buildJava {

	export  COMPONENT=$1
	export  JAVAPATH=${PROJECT_PATH}/${COMPONENT}

	echo ">>> Building component \"${COMPONENT}\" at ${JAVAPATH}"
	echo

	#	build the component

	echo "> Building source code"

	cd ${JAVAPATH}
	mvn clean install
}

#	build all project components

echo ">>>>> Building GoLang components"
echo

buildGoLang "request"
buildGoLang "requestLoader"
buildGoLang "project"

echo ">>>>> Building Java components"
echo

buildJava ebcdic-spring-boot/complete

