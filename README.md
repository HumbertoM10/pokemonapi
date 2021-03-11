# pokemonapi
*This is a retake on the pokeapi*

## Running the project
#### local
Run the following commands
    $ git clone https://github.com/HumbertoM10/pokemonapi.git
    $ go run main.go

This will be running in http://localhost on port :3000

#### docker
Run the following commands
    $ git clone https://github.com/HumbertoM10/pokemonapi.git
    $ docker build . --tag <tag-name>
    $ docker run -d -p <desired-port>:3000 <tag-name>

This will be running in http://localhost on port :3000

## Deployment
Using docker the project already has an image uploaded to humbertoe10/pokeapi

This project is also deployed in a AWS EKS using Kubernetes as the container orchestrator.

The configuration used for this project is stored in the following file of the project: k8s/eksctl.yaml
This service can be launched by the following command:
    $ eksctl create cluster -f <file>

The Kubernetes deployment was made using this file: k8s/deployment.yaml
This can be added to the cluster using the following command
    $ kubectl apply -f <file>

The Kubernetes service was made using this file: k8s/service.yaml
This can be added to the cluster using the following command
    $ kubectl apply -f <file>
       
## Continuos Integration and Delivery
Using GitHub Actions the flow is the following:
Test: I used the following GitHub Action from Setup Go environment
https://github.com/marketplace/actions/setup-go-environment

Build: I used the following GitHub Action from Build and push Docker images
https://github.com/marketplace/actions/build-and-push-docker-images

Deploy: I used the following GitHub Action from kubectl-aws-eks
https://github.com/marketplace/actions/kubectl-aws-eks

## Documentation
- [handler](https://github.com/HumbertoM10/pokemonapi/packages/handler)
- [parser](https://github.com/HumbertoM10/pokemonapi/packages/parser)

## RESTApi Documentation
**Advantage**(endpoint)
Advantage is an endpoint that stores the data of 2 given pokemons telling which one has an advantage over the other with the following data:
	- HasAdvantage:	Does the first pokemon has an advantage over the second one? (true or false)
	- DmgTaken:		Multiplier of damage recived by the first pokemon from the second pokemon
	- DmgDone:		Multiplier of damage done by the first pokemon to the second pokemon
	- Poke1:		Name of the first pokemon
	- Poke2:		Name of the second pokemon
	- Explanation:	Explanation on why the first pokemon has an advantage or not over the second pokemon

**CommonMoves**(endpoint)
commonMoves is an endpoint that stores the data of all given pokemons telling which moves are the ones they have in common:
	- Language:			Code of the language in which the data is going to be stored
	- Pokemons:			Data of the pokemon which moves were compared
	- MovesInCommon:	A list of all the moves they have in common