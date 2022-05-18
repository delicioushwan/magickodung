# maginkodung

마법고둥

local 실행
minikube설치후

minikube start

eval $(minikube -p minikube docker-env)

kubectl apply -f 를 이용하여 db front back 생성

minikube라 nodeport등을 간단하게 할수없음 ingress는 아직 설정이 되지않음
port-forwarding을 이용하여
back : 9091:9091
front : 3000:3000
db: 3306:9876

으로 로컬 port를 deployment와 forwarding한뒤 브라우저를 통해 접속 가능
