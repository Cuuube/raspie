#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>

#define PIN_BTN 0
#define LOW  0
#define HIGH 1


int main(int argc, char const *argv[])
{
    printf("----- START -----\n");
    int currentStatus = OUTPUT;

    if (wiringPiSetup() == -1) {
        printf("setup wiringPi failed!");
        return -1;
    }

    pinMode(PIN_BTN, INPUT);

    while(1)
    {
        delay(10);

        // 以数字电路模式，读取当前针脚的高或低
        int status = digitalRead(PIN_BTN);
        // printf("Read status is %d, current status is %d !\n", status, currentStatus);

        // 如果针脚接收到信号有变化，则打印log
        if (status == LOW && currentStatus != status)
        {
            printf("Btn down!\n");
            currentStatus = status;
        }
        else if (status == HIGH && currentStatus != status)
        {
            printf("Btn up!\n");
            currentStatus = status;
        }
    }
    printf("----- END -----\n");
    return 0;
}
