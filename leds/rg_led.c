#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>

// which gives you at least the -127 to 127 range. (-128 to 127 is common)
// unsigned char, which gives you at least the 0 to 255 range.
#define uchar unsigned char
#define PIN_RED 1   // 红色信号针脚
#define PIN_GREEN 2 // 绿灯信号针脚
#define COLOR_RANGE 0xff // 颜色范围

int init() {
    // 初始化
    if (wiringPiSetup() == -1) { // 初始化pi
        printf("setup wiringPi failed!");
        return -1;
    }

    // 定义pinMode是输出信号（比如小灯），还是接收信号（比如开关）
    pinMode(PIN_RED, OUTPUT);
    pinMode(PIN_GREEN, OUTPUT);

    // 设置可接受的范围值
    softPwmCreate(PIN_RED, 0x00, COLOR_RANGE);
    softPwmCreate(PIN_GREEN, 0x00, COLOR_RANGE);
    return 0;
}

void setColor(uchar redValue, uchar greenValue) {
    // 设置电平，必须要在之前的范围值里
    if (redValue < 0 || redValue > COLOR_RANGE) {
        return;
    }
    if (greenValue < 0 || greenValue > COLOR_RANGE) {
        return;
    }
    softPwmWrite(PIN_RED, COLOR_RANGE - redValue);
    softPwmWrite(PIN_GREEN, COLOR_RANGE - greenValue);
}

int main(int argc, char const *argv[])
{
    int initResult = init();
    if (initResult != 0)
    {
        return -1;
    }

    printf("----- START -----\n");

    setColor(0xff, 0x00); // 红
    delay(2000);

    setColor(0x00, 0xff); // 绿
    delay(2000);

    // setColor(0x00, 0x00); // 本元件不可以同时发两种光
    setColor(0xff, 0x00); // 红
    delay(2000);

    setColor(0x00, 0x80); // 半绿
    delay(2000);

    setColor(0x80, 0x00); // 半红
    delay(2000);

    setColor(0xff, 0xff); // 重置
    delay(2000);
    printf("----- END -----\n");
    // 越低反而颜色权重大
    return 0;
}
