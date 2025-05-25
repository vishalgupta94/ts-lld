class Light {
    turnOff(){
        console.log("TURN OFF")
    }

    turnOn(){
        console.log("TURN OFF")
    }
}

interface Command {
    execute():void
    undo():void
}

class TurnOnCommannd implements Command{
    private light: Light
    constructor(light: Light){
        this.light = light
    }
    undo(){
        this.light.turnOff()
    }
    execute(){
        this.light.turnOn()
    }
}

class TurnOffCommannd implements Command{
    private light: Light
    constructor(light: Light){
        this.light = light
    }
    execute(){
        this.light.turnOff()
    }
    undo(){
        this.light.turnOn()
    }
}