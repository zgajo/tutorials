import { ColorHelper } from "./src/ColorHelper";

// index.js
// I've combined sketch.js and Star.js here to make
// it easier to share variable speed
import p5 from "p5";
const stars: Array<Star> = [];
let speed = 1;

const myp5 = new p5(() => {});

myp5.setup = function setup() {
  myp5.createCanvas(600, 600);
  for (let i = 0; i < 800; i++) {
    stars[i] = new Star();
  }
};

myp5.draw = function draw() {
  speed = myp5.map(myp5.mouseX, 0, myp5.width, 0, 50);
  myp5.background(0);
  myp5.translate(myp5.width / 2, myp5.height / 2);
  for (let i = 0; i < stars.length; i++) {
    stars[i].update();
    stars[i].show();
  }
};

class Star {
  x: any = null;
  y: any = null;
  z: any = null;
  pz: any = null;
  constructor() {
    this.x = myp5.random(-myp5.width, myp5.width);
    this.y = myp5.random(-myp5.height, myp5.height);
    this.z = myp5.random(myp5.width);
    this.pz = this.z;
  }

  update() {
    this.z = this.z - speed;
    if (this.z < 1) {
      this.z = myp5.width;
      this.x = myp5.random(-myp5.width, myp5.width);
      this.y = myp5.random(-myp5.height, myp5.height);
      this.pz = this.z;
    }
  }

  show() {
    myp5.fill(255);
    myp5.noStroke();

    var sx = myp5.map(this.x / this.z, 0, 1, 0, myp5.width);
    var sy = myp5.map(this.y / this.z, 0, 1, 0, myp5.height);

    var r = myp5.map(this.z, 0, myp5.width, 16, 0);
    myp5.ellipse(sx, sy, r, r);

    var px = myp5.map(this.x / this.pz, 0, 1, 0, myp5.width);
    var py = myp5.map(this.y / this.pz, 0, 1, 0, myp5.height);

    this.pz = this.z;

    myp5.stroke(255);
    myp5.line(px, py, sx, sy);
  }
}

export const myp5_2 = new p5(() => {});

// GLOBAL VARS & TYPES
let numberOfShapes = 15;
let slider_speed: p5.Element;

// P5 WILL AUTOMATICALLY USE GLOBAL MODE IF A DRAW() FUNCTION IS DEFINED
myp5_2.setup = function () {
  console.log("🚀 - Setup initialized - P5 is running");

  // FULLSCREEN CANVAS
  myp5_2.createCanvas(myp5_2.windowWidth, myp5_2.windowHeight);

  // SETUP SOME OPTIONS
  myp5_2.rectMode(myp5_2.CENTER).noFill().frameRate(30);

  // SPEED SLIDER
  slider_speed = myp5_2.createSlider(0, 15, 3, 1);
  slider_speed.position(10, 620);
  slider_speed.style("width", "80px");
};

// p5 WILL HANDLE REQUESTING ANIMATION FRAMES FROM THE BROWSER AND WIL RUN DRAW() EACH ANIMATION FROME
myp5_2.draw = function () {
  // CLEAR BACKGROUND
  myp5_2.background(0);
  // TRANSLATE TO CENTER OF SCREEN
  myp5_2.translate(myp5_2.width / 2, myp5_2.height / 2);

  const colorsArr = ColorHelper.getColorsArray(numberOfShapes);
  const baseSpeed = (myp5_2.frameCount / 500) * <number>slider_speed.value();
  for (var i = 0; i < numberOfShapes; i++) {
    const npoints = 3 + i;
    const radius = 20 * i;
    const angle = myp5_2.TWO_PI / npoints;
    const spin = baseSpeed * (numberOfShapes - i);

    myp5_2.strokeWeight(3 + i).stroke(colorsArr[i]);

    myp5_2.push();
    myp5_2.rotate(spin);
    // DRAW
    myp5_2.beginShape();
    for (let a = 0; a < myp5_2.TWO_PI; a += angle) {
      let sx = myp5_2.cos(a) * radius;
      let sy = myp5_2.sin(a) * radius;
      myp5_2.vertex(sx, sy);
    }
    myp5_2.endShape(myp5_2.CLOSE);
    // END:DRAW
    myp5_2.pop();
  }
};

// p5 WILL AUTO RUN THIS FUNCTION IF THE BROWSER WINDOW SIZE CHANGES
myp5_2.windowResized = function () {
  myp5_2.createCanvas(myp5_2.windowWidth, myp5_2.windowHeight);
};
