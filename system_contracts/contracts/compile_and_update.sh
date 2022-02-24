dst_json_file=$1

if [ ! -f "$1" ];then
  echo "genesis.json文件目录错误: $1"
  exit
fi

if [ `solc --version | grep "0.6.4" -c` != 1 ];then
  echo "solc version wrong. use 0.6.4 for now"
  exit
fi

#### add contract info if here
file_list=("CrossChain.sol" "TokenExchange.sol")
address_list=("0000000000000000000000000000000000000100" "0000000000000000000000000000000000000200")
res_list=("" "")
####

length=${#file_list[*]}

if [ ${#file_list[*]} != ${#address_list[*]} ] 
then
    echo "!!!! address list and contract file list not equal !!!!"
    exit
fi

for i in `seq $length`
do
  index=i-1
  contract_file=${file_list[index]}
  address=${address_list[index]}
  bin_runtime_data="$( solc --bin-runtime $contract_file | grep "$contract_file" -A2 | tail -n 1)"
  abi_data="$( solc --abi $contract_file | grep "$contract_file" -A2 | tail -n 1)"
  if [ ${#abi_data} == 0 -o ${#bin_runtime_data} == 0 ] 
  then
    echo "length: ${#abi_data}"
    echo "length: ${#bin_runtime_data}"
    echo "!!!! compile contract $contract_file failed !!!!"
    exit
  else
    ## debug info
    # echo "address: $address"
    # echo "bin-runtime: $bin_runtime_data"
    # echo "abi: $abi_data"
    
    # replace contract binary code
    old_code="$(  grep "\"$address\"" $dst_json_file -A5 | grep "code" )"
    res="$(sed -i "s/$old_code/            \"code\": \"0x$bin_runtime_data\"/g" $dst_json_file)"

    # save contract object code used in console.
    short_name=`expr substr $contract_file 1 6`
    res_list[index]="var $short_name = eth.contract($abi_data).at($address)"
  fi
done

echo "contract object code used in console"
for i in `seq $length`
do
  index=i-1
  echo "=====${file_list[index]}====="
  echo "${res_list[index]}"
  echo "========================"
  echo ""
done