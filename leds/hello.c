// #include <wiringPi.h>
// #include <softPwm.h>
#include <stdio.h>

// which gives you at least the -127 to 127 range. (-128 to 127 is common)
// unsigned char, which gives you at least the 0 to 255 range.
#define uchar unsigned char
#define PIN_RED 0
#define PIN_GREEN 1


int main(int argc, char const *argv[])
{
    printf("----- START -----\n");

    // delay(500);
    int a = 500;
    int b = 0xff;
    uchar c = a;
    uchar d = b;
    printf("%d\n", a);
    printf("%d\n", b);
    printf("%d\n", c);
    printf("%d\n", d);

    printf("----- EMD -----\n");
    /* code */
    return 0;
}


