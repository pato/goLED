#include <PololuLedStrip.h>

// Create an ledStrip object and specify the pin it will use.
PololuLedStrip<12> ledStrip;

// Create a buffer for holding the colors (3 bytes per color).
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
  
void setup()
{
  // Start up the serial port, for communication with the PC.
  Serial.begin(115200);
  Serial.println("Ready to receive colors."); 
}

void simplerSerial() {
 char c = Serial.peek();
    if (!(c >= '0' && c <= '9'))
    {
      Serial.read();
    }else{
      // Read the color from the computer.
      rgb_color color;
      color.red = Serial.parseInt();
      color.green = Serial.parseInt();
      color.blue = Serial.parseInt();
      
      uint8_t index = Serial.parseInt();
      
      colors[index] = color;
      ledStrip.write(colors, LED_COUNT);
    } 
}

void simpleSerial() {
  union stream_data data;
  Serial.readBytes((char*)&data, 4);
  Serial.print(data.rgbd.r);Serial.print(", ");
  Serial.print(data.rgbd.g);Serial.print(", ");
  Serial.print(data.rgbd.b);Serial.print(", ");
  Serial.print(data.rgbd.index);
  Serial.println();
  colors[data.rgbd.index] = data.rgb;
  ledStrip.write(colors, LED_COUNT);
  
  //rgb_color rgb = (rgb_color) rgbd;
  
  //colors[rgbd.index] = ((rgb_color) rgbd);
}

void stripSerial() {
  Serial.readBytes((char*)colors, LED_COUNT * 3);
  ledStrip.write(colors, LED_COUNT);
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

void loop()
{
  if (Serial.available()){
    improvedSerial();
  }
}
