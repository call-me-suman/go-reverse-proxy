package configs


import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type  Resource struct {
	Name string
	Destination_url string
	Endpoint string
}



type  Config struct{
	Server struct {
		Host string
		Listen_port string
	 }
	Resources []Resource
	
	Rate struct{
		Reqs int
		Persec int
	}

}


var config *Config
func NewConfiguration()(*Config , error){
	viper.AddConfigPath("data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))
	

	err:= viper.ReadInConfig()

	if err != nil{
		return nil,fmt.Errorf("Error loading the config file :%s",err)
	}

	err = viper.Unmarshal(&config)

	if err != nil{
return nil, fmt.Errorf("Error reading the config File :%s",err)
	}

	return config,nil






}