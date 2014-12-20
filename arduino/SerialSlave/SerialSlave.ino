#include <PololuLedStrip.h>

PololuLedStrip<12> ledStrip;

#define LED_COUNT 60
rgb_color colors[LED_COUNT];

typedef struct rgb_color_data{
  unsigned char r, g, b;
  uint8_t index;
} rgb_color_data;

union stream_data {
  rgb_color rgb;
  rgb_color_data rgbd;
};


void setup() {
  Serial.begin(115200);
  Serial.println("Ready to receive colors."); 
}

void improvedSerial() {
  uint8_t header = Serial.read();
  if (header == 'f') {
    ledStrip.write(colors, LED_COUNT);
  } else if (header == 's') {
    union stream_data data;
    Serial.readBytes((char*)&data, 4);
    colors[data.rgbd.index] = data.rgb;
  } else if (header == 'c') {
    for (int i = 0; i < LED_COUNT; i++) {
       colors[i] = (rgb_color){0,0,0}; 
    }
    ledStrip.write(colors, LED_COUNT);
  }
}

void loop() {
  if (Serial.available()){
    improvedSerial();
  }
}
