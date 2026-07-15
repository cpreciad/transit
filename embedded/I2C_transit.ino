/*
  Author: Carlo Preciado 
  Date:  July 2026
*/

#include <Bonezegei_LCD1602_I2C.h>
#include <time.h>

bool availableWithTimeout(unsigned long);
void renderScreen();
void sleepScreen();

Bonezegei_LCD1602_I2C lcd(0x27);

char buf[BUFSIZ];
char DELIM = '\n';
unsigned long SERIAL_TIMEOUT_SEC = 60;

void setup() {
  // begin code
  lcd.begin();
  Serial.begin(9600);  
  
  lcd.print("Carlo Preciado");
  lcd.setPosition(0, 1);      //param1 = X   param2 = Y
  lcd.print("Good Morning!");
  delay(2000);
}

void loop() {
  bool isAvailable = availableWithTimeout(SERIAL_TIMEOUT_SEC);
  if (isAvailable){
    renderScreen();
  }
  else{
    sleepScreen();
  }
}

// effectively spins on Serial.available() until a given timeout
// returns true if serial became available
// returns false if serial timed out before data was sent
bool availableWithTimeout(unsigned long timeout){
  unsigned long start_time = millis();
  
  while(Serial.available() == 0) {
    unsigned long current_time = millis() - start_time;
    if (current_time >= timeout * 1000){
      return false;
    }
  }
  return true;
}

// these two just abstract the ldc calls

void renderScreen(){
  // make sure backlight is on
  lcd.setBacklight(1);

  int renderRow = 0;
  lcd.clear();
  for (String str = Serial.readStringUntil(DELIM); str != NULL; str = Serial.readStringUntil(DELIM)){
    lcd.setPosition(0, renderRow);
    lcd.print(str.c_str());
    renderRow++;
  }
}

void sleepScreen(){
  char buf[BUFSIZ];
  lcd.clear();
  lcd.setPosition(0,0);
  sprintf(buf, "timeout: %lu s", SERIAL_TIMEOUT_SEC);
  lcd.print(buf);
  lcd.setPosition(0, 1);
  lcd.print("time to sleep...");
  delay(4000);
  lcd.clear();
  lcd.setBacklight(0);
}
