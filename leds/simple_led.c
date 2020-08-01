#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>

// which gives you at least the -127 to 127 range. (-128 to 127 is common)
// unsigned char, which gives you at least the 0 to 255 range.
#define uchar unsigned char
#define PIN_RED 0
#define PIN_GREEN 1

void init() {
    // 初始化
    softPwmCreate(PIN_RED, 0, 100);
    softPwmCreate(PIN_GREEN, 0, 100);
}

void setColor(uchar redValue, uchar greenValue) {
    softPwmWrite(PIN_RED, redValue);
    softPwmWrite(PIN_GREEN, greenValue);
}

int main(int argc, char const *argv[])
{
    printf("----- START -----");

    setColor(0xff, 0x00);
    delay(500);

    setColor(0x00, 0xff);
    delay(500);

    setColor(0xff, 0xff);
    delay(500);

    setColor(0x80, 0x80);
    delay(500);

    setColor(0x80, 0x80);
    delay(500);

    printf("----- END -----");
    /* code */
    return 0;
}


