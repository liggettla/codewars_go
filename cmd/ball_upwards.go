/*
You throw a ball vertically upwards with an initial speed v (in km per hour). The height h of the ball at each time t is given by h = v*t - 0.5*g*t*t where g is Earth's gravity (g ~ 9.81 m/s**2). A device is recording at every tenth of second the height of the ball. For example with v = 15 km/h the device gets something of the following form: (0, 0.0), (1, 0.367...), (2, 0.637...), (3, 0.808...), (4, 0.881..) ... where the first number is the time in tenth of second and the second number the
height in meter.

Task
Write a function max_ball with parameter v (in km per hour) that returns the time in tenth of second of the maximum height recorded by the device.

Examples:
max_ball(15) should return 4

max_ball(25) should return 7

Notes
Remember to convert the velocity from km/h to m/s or from m/s in km/h when necessary.
The maximum height recorded by the device is not necessarily the maximum height reached by the ball.
*/

package main

import("fmt")

// h = v*t - 0.5*g*t*t
// v = (h + 0.5*g*t*t) / t
// Vy = V0 sin(θ) - g t

// convert km/hr to m/s
func meter_second(velocity int) float64 {
    converted_velocity := float64(velocity)
    converted_velocity = converted_velocity / 360 // seconds
    converted_velocity = converted_velocity * 1000 // meters
    return converted_velocity
}

func time_of_flight(v float64) float64 {
    var g float64 = 9.8
    var time float64 = 0
    var t float64 = 0.1

    // time ball is going up
    for v>= 0 {
        time += t
        v = v - g*t
        fmt.Println(v,t)
    }
    return time-0.1
}

func main() {
    var velocity int = 25 // in km/hr

    converted_velocity := meter_second(velocity)
    tof := time_of_flight(converted_velocity)
    fmt.Println(tof)
    fmt.Println(0.1*70)
}
