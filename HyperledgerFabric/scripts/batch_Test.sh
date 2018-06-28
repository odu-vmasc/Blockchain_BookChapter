#!/bin/bash

# Importing useful functions for cc testing
if [ -f ./func.sh ]; then
 source ./func.sh
elif [ -f scripts/func.sh ]; then
 source scripts/func.sh
fi

tot=1
val=$1
iter=1
end=1

ids = 
idcount=1;
instantiateMult 0 bank asset
for t in {1..500}
do
	idcount=$((idcount+1))
	id1=a$idcount

	idcount=$((idcount+1))
	id2=a$idcount
	instantiateMult 0 $id1 $id2
	echo_b "==============0 "$id1" "$id2"========="
done

fulltime=$(date +%s)
peerid=0
#while [ $iter -lt $tot ]; do
for i in {1..5}
do
	echo_b "Channel name : "$CHANNEL_NAME
	
	echo_b "====================Query the existing value of a===================================="
	chaincodeQuery 0 $val
	
	for q in {2..1000}
	do
		echo_b "=====================Invoke a transaction to transfer 1 from a to b=================="
		chaincodeInvokeMult $peerid a$q bank
		val=$((val - 1))
		sleep 5 
		
		if [$peerid -gt 3]
		then
			peerid=0
		else
			peerid=$(peerid+1)
		fi
	done
	echo_b "=====================Check if the result of a is $val==================================="
	#chaincodeCheckMult 0 a100 $val	
	iter=$((iter + 1))
done
#done
echo "testcomplete,$(($(date +%s)-fulltime))" >> report.txt
#echo "$cf,$2,$(($(date +%s)-starttime))" >> report.txt
echo
echo_g "=====================All GOOD, MVE Test completed ===================== "
echo
exit 0
