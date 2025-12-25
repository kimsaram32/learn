#include <Servo.h>

// pins
const int PIN_SOUND = A1;
const int PIN_DISTANCE = A5;
const int PIN_TEMP = A3;

const int PIN_LED_RED = 2;
const int PIN_LED_GREEN = 3;
const int PIN_LED_BLUE = 4;
const int PIN_LED_YELLOW = 5;
const int PIN_DC = 6;
const int PIN_BUTTON_RED = 8;
const int PIN_BUTTON_BLUE = 9;
const int PIN_BUTTON_YELLOW = 10;
const int PIN_BUTTON_GREEN = 11;

// door
const int door_distance_threshold = 500;
const int door_start_angle = 0;
const int door_end_angle = 180;
const int door_close_button_pin = PIN_BUTTON_RED;
int door_current_angle = 0;
int door_opened = 0;

Servo servo;

void setup(){
  // door
  servo.attach(PIN_DC);
  pinMode(door_close_button_pin, INPUT);

  Serial.begin(9600);
}

void loop(){
  int distance = analogRead(PIN_DISTANCE);

  Serial.print("Distance is : "); 
  Serial.print(distance); 
  Serial.print("\n");

  if (distance < door_distance_threshold && !door_opened) {
    door_opened = 1;
    while (door_current_angle < door_end_angle) {
      servo.write(++door_current_angle);
      delay(5);
    }
  }

  int door_close_button_state = digitalRead(door_close_button_pin);
  if (door_close_button_state == HIGH && door_opened) {
    door_opened = 0;
    while (door_current_angle > door_start_angle) {
      servo.write(--door_current_angle);
      delay(5);
    }
  }
  delay(500); 
}
