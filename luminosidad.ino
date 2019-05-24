#include <ArduinoHttpClient.h>
#include <ArduinoJson.h>


//#include "Servo.h"

#include <ESP8266HTTPClient.h>
#include <ESP8266HTTPUpdateServer.h>
#include <ESP8266WiFi.h>
//#include <WiFiClient.h>

//Ultrasonicos de salida

#define foto  A0//0//D8

int valor=0;
const char* ssid = "AP";
const char* password = "12345678";
const char* server = "http://192.168.0.106:2000/boton";

//const char* ssid = "Mantenimiento3";
//const char* password = "lucero2808a";
//const char* server = "http://3.16.165.22:3000/hidroponia";

//const char* ssid = "nerdstudio";
//const char* password = "vvcdk1dN";
//const char* server = "http://10.42.0.19:800/boton";

String ID="NodeMCU Lumino";

void setup() {
  Serial.begin(9600);
   //Inicializando WiFi
  Serial.print("Conectando ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");}
  Serial.println("WiFi conectado");
  pinMode(foto, INPUT);
}


void loop() {
  valor=analogRead(foto);
  //valor=1;
  post(ID, valor);
  delay(100);
}


void post (String id, int valor ){
  HTTPClient http;
  String json;
  StaticJsonBuffer<800> jsonBuffer;
  String valor2=(String)valor;
  JsonObject& root = jsonBuffer.createObject();
  root["ID"] = id;
  
  root["LUMINO"]= valor2;
  //root["NOMBRE"] = valor2;
  root.printTo(json); 
  http.begin(server);
  http.addHeader("Content-Type", "application/json");
  http.POST(json);
  http.writeToStream(&Serial);
  delay(100);
  http.end();
  }
