class Wheel {
    constructor(id, x, y, r, userEmail, errorCallback) {
        if (typeof id !== "string" || id.length === 0) {
            throw "id has invalid type"
        }
        if (typeof userEmail !== "string" || userEmail.length === 0) {
            throw "userEmail has invalid type"
        }
        if (typeof x !== "number") {
            throw "x has invalid type"
        }
        if (typeof y !== "number") {
            throw "y has invalid type"
        }
        if (typeof r !== "number") {
            throw "r has invalid type"
        }
        if (typeof errorCallback !== "function") {
            throw "errorCallback has invalid type"
        }

        this.id = id
        this.x = x
        this.y = y
        this.r = r
        this.userEmail = userEmail
        this.errorCallback = errorCallback

        this.genericParsingErrMsg = "Failed to parse the server's response, if the issue persists, contact the customer support."

        this.container = document.getElementById(id)
        if (this.container === null) {
            throw `element of id '${id}' not found`
        }
        this.container.style.cursor = "pointer"

        this.canvas2 = document.createElement("canvas")
        this.canvas2.style.width = "100%"
        this.canvas2.style.height = "100%"
        this.canvas2.style.position = "absolute"
        this.canvas2.style.top = "0pxs"
        this.canvas2.width = window.innerWidth
        this.canvas2.height = window.innerHeight
        this.ctx2 = this.canvas2.getContext("2d")

        this.canvas = document.createElement("canvas")
        this.canvas.style.width = "100%"
        this.canvas.style.height = "100%"
        this.canvas.style.position = "absolute"
        this.canvas.style.top = "0pxs"
        this.canvas.width = window.innerWidth
        this.canvas.height = window.innerHeight

        this.ctx = this.canvas.getContext("2d")

        this.rotAngle = 0
        this.rotSpeed = 1
        this.spin = false
        this.resolution = 1

        this.draw()

        this.initSpinAnimation()

        this.initSpinListener()
        
        this.container.appendChild(this.canvas2)
        this.container.appendChild(this.canvas)
    }

    draw() {
        // Primary shape.
        this.ctx.beginPath()
        this.ctx.arc(this.x, this.y, this.r, 0, 2 * Math.PI)
        this.ctx.strokeStyle = "#8BE9FD"
        this.ctx.stroke()
        this.ctx.fillStyle = "#44475A"
        this.ctx.fill()

        const lines = 12
        this.angle = 360/lines

        const initX = this.x
        const initY = this.y

        const posOffset = this.r

        const prevR = this.angle*(lines-1) * (Math.PI/180)
        let prevX = initX + Math.cos(prevR)*posOffset
        let prevY = initX + Math.sin(prevR)*posOffset

        let p1x, p1y, p2x, p2y = 0

        for (let i = 0; i < 12; i++) {
            const r = this.angle*i * (Math.PI/180)
            const x = initX + Math.cos(r)*posOffset
            const y = initX + Math.sin(r)*posOffset
        
            this.ctx.beginPath()
            this.ctx.lineTo(initX, initY)
            this.ctx.lineTo(x, y-initY)
            this.ctx.stroke()

            this.ctx.beginPath()
            this.ctx.fillStyle = "#FF5555"
            if (i === 0 || i === 3 || i === 7 || i === 10) {
                this.ctx.fillStyle = "#50FA7B"
            }
            this.ctx.font = "22px arial"
            this.ctx.fillText(`${i+1}`, x-(prevX-x)/2, y-initY-(prevY-y)/2)
        
            prevX = x
            prevY = y

            if (i === 0) {
                p1x = x-(prevX-x)/2
                p1y = y-initY-(prevY-y)/2
            }
            if (i === 1) {
                p2x = x-(prevX-x)/2
                p2y = y-initY-(prevY-y)/2
            }
        }

        this.ctx2.beginPath()
        this.ctx2.lineTo(p2x, p2y-(p2y-p1y)/2)
        this.ctx2.lineTo(window.innerWidth, window.innerHeight)
        this.ctx2.strokeStyle = "#50FA7B"
        this.ctx2.stroke()
        this.ctx2.fill()
        this.ctx2.beginPath()

        // Inner circle.
        this.ctx.beginPath()
        this.ctx.arc(this.x, this.y, this.r/16, 0, 2 * Math.PI)
        this.ctx.fillStyle = "#8BE9FD"
        this.ctx.fill()
    }

    initSpinListener() {
        this.canvas.addEventListener('click', () => {
            if (this.spin === true) {
                return
            }
            this.spin = true
            setTimeout(() => {
                const req = new XMLHttpRequest()
                req.open("GET", `http://127.0.0.1:8081/v1/spin/${this.userEmail}`, true)

                req.onload = () => {
                    if (req.status !== 200) {
                        this.spin = false
                        if (req.responseText.includes("not enough credits")) {
                            this.errorCallback(`Try next time`)
                        }
                        this.errorCallback(`Request failed to get spin details. Status code ${req.status}: ${req.response}`)
                        return
                    }

                    try {
                        const resp = JSON.parse(req.response)
                        if (
                            typeof resp.SpinID !== "string" ||
                            typeof resp.Number !== "number" ||
                            typeof resp.Win !== "boolean"
                        ) {
                            this.errorCallback(this.genericParsingErrMsg)
                            console.error(resp)
                            return
                        }

                        this.selectedNumber = resp.Number
                        this.win = resp.Win
                        this.lastSpinID = resp.SpinID

                        console.log(resp)
                    }
                    catch(e) {
                        this.errorCallback(this.genericParsingErrMsg)
                        console.error(e)
                        return
                    }
                }

                req.send()
            }, 1000)
        }, false)
    }

    initSpinAnimation() {
        setInterval(() => {
            if (this.spin === false) {
                return
            }
            
            this.rotAngle == 360
                ? this.rotAngle = 0
                : this.rotAngle += this.rotSpeed
            
            this.canvas.style.setProperty('transform', `rotate(${this.rotAngle}deg)`)
        
            if (this.selectedNumber > 0 && this.rotAngle === (12-this.selectedNumber /* +1 since the server indexes from 0 */)*this.angle) {
                this.spin = false
                this.selectedNumber = -1

                console.log(this.win)
                if (this.win) {
                    const req = new XMLHttpRequest()
                    req.open("GET", `http://127.0.0.1:8081/v1/prizes/${this.lastSpinID}`, true)

                    req.onload = () => {
                        try {
                            const resp = JSON.parse(req.response)
                            if (typeof resp.Prize !== "number") {
                                this.errorCallback(this.genericParsingErrMsg)
                                console.error(resp)
                                return
                            }
    
                            this.errorCallback(`You won ${resp.Prize} points!!!`)
                        }
                        catch(e) {
                            this.errorCallback(this.genericParsingErrMsg)
                            console.error(e)
                            return
                        }
                    }

                    req.send()
                } else {
                    // this.errorCallback(`Try again...`)
                }

                return
            }
        }, this.resolution)
    }
}