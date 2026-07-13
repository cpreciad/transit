/*
  Hello World 
  Author: Bonezegei (Jofel Batutay)
  Date:  January 2024
*/

#include <Bonezegei_LCD1602_I2C.h>

Bonezegei_LCD1602_I2C lcd(0x27);

char buf[BUFSIZ];

void setup() {
  lcd.begin();
  lcd.print("Carlo Preciado");
  lcd.setPosition(0, 1);      //param1 = X   param2 = Y
  lcd.print("Good Morning!");
  delay(2000);
  lcd.setPosition(0, 0);
  Serial.begin(9600);
}

void loop() {
  char *token;
  char *saveptr; 
  char *delim = "\n";

  while(Serial.available() > 0){
    String str = Serial.readString();
    str.toCharArray(buf, BUFSIZ);
    lcd.clear();
    
    // first part of message
    lcd.setPosition(0, 0);
    token = strtok_r(buf, delim, &saveptr);
    lcd.print(token);

    // second part, should contain times
    lcd.setPosition(0, 1);
    token = strtok_r(NULL, delim, &saveptr);
    lcd.print(token);

    delay(500);

  }
  
}


