#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>

#define PIN_BTN 0


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
        int status = digitalRead(PIN_BTN);
        // printf("Read status is %d, current status is %d !\n", status, currentStatus);
        

        if (status == INPUT && currentStatus != status)
        {
            printf("Btn down!\n");
            currentStatus = status;
        }
        else if (status == OUTPUT && currentStatus != status)
        {
            printf("Btn up!\n");
            currentStatus = status;
        }
    }
    printf("----- END -----\n");
    return 0;
}
