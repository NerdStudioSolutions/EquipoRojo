#include <ArduinoHttpClient.h>
#include <ArduinoJson.h>


//#include "Servo.h"

#include <ESP8266HTTPClient.h>
#include <ESP8266HTTPUpdateServer.h>
#include <ESP8266WiFi.h>
//#include <WiFiClient.h>

#include "DHT.h"        // including the library of DHT11 temperature and humidity sensor
#define DHTTYPE DHT11   // DHT 11
#define sensor A0

int ph=0;
int humedad=0;
int temp=0;
const char* ssid = "AP";
const char* password = "12345678";
const char* server = "http://192.168.0.106:2000/boton";
//const char* server = "http://10.42.0.19:800/boton";
String ID="NodeMCU dht11";


#define dht_dpin 0
DHT dht(dht_dpin, DHTTYPE); 
void setup(void)
{ 
   Serial.begin(9600);
   //Inicializando WiFi
  Serial.print("Conectando ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");}
  Serial.println("WiFi conectado");
  pinMode(dht_dpin, INPUT);
  pinMode(sensor, INPUT);

  
  dht.begin();
  Serial.begin(9600);
  Serial.println("Humidity and temperature\n\n");
  delay(700);

}
void loop() {
    float humedad = dht.readHumidity();
    float temp = dht.readTemperature();         
    Serial.print("Current humidity = ");
    Serial.print(humedad);
    Serial.print("%  ");
    Serial.print("temperature = ");
    Serial.print(temp); 
    Serial.println("C  ");
  delay(800);



 int ph = map(analogRead(sensor), 0, 1023, 100, 0);
  
  Serial.print("Humedad: ");
  Serial.print(ph);
  Serial.println("%");



   post(ID, humedad,temp,ph);
  delay(100);
}


void post (String id, int temp, int humedad, int ph ){
  HTTPClient http;
  String json;
  StaticJsonBuffer<800> jsonBuffer;
  JsonObject& root = jsonBuffer.createObject();
  String temp2=(String) temp;
  String humedad2=(String) humedad;
  String ph2=(String) ph;
  root["ID"] = id;
  root["TEMP"] = temp2;
  root["HUMEDAD"] = humedad2;
  root["PH"] = ph2;
  root.printTo(json); 
  http.begin(server);
  http.addHeader("Content-Type", "application/json");
  http.POST(json);
  http.writeToStream(&Serial);
  delay(100);
  http.end();
  }
