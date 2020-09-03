#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>

#define PIN_BTN 7


int main(int argc, char const *argv[])
{
    printf("----- START -----");

    if (wiringPiSetup() == -1) {
        print("setup wiringPi failed!");
        return -1;
    }

    pinMode(PIN_BTN, INPUT);

    while(1)
    {
        if (digitalRead(PIN_BTN) == INPUT)
        {
            delay(10);
            print('Btn down!');
        }
        else if (digitalRead(PIN_BTN) == OUTPUT)
        {
            print('Btn up!');
        }
    }
    printf("----- END -----");
    return 0;
}


