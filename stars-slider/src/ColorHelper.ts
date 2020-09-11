import { myp5_2 } from "..";

export class ColorHelper {
  private static getColorVector(c: p5.Color) {
    return myp5_2.createVector(myp5_2.red(c), myp5_2.green(c), myp5_2.blue(c));
  }

  public static rainbowColorBase() {
    return [
      myp5_2.color("red"),
      myp5_2.color("orange"),
      myp5_2.color("yellow"),
      myp5_2.color("green"),
      myp5_2.color(38, 58, 150), // blue
      myp5_2.color("indigo"),
      myp5_2.color("violet"),
    ];
  }

  public static getColorsArray(
    total: number,
    baseColorArray: p5.Color[] = null
  ): p5.Color[] {
    if (baseColorArray == null) {
      baseColorArray = ColorHelper.rainbowColorBase();
    }
    var rainbowColors = baseColorArray.map((x) => this.getColorVector(x));

    let colours = new Array<p5.Color>();
    for (var i = 0; i < total; i++) {
      var colorPosition = i / total;
      var scaledColorPosition = colorPosition * (rainbowColors.length - 1);

      var colorIndex = Math.floor(scaledColorPosition);
      var colorPercentage = scaledColorPosition - colorIndex;

      var nameColor = this.getColorByPercentage(
        rainbowColors[colorIndex],
        rainbowColors[colorIndex + 1],
        colorPercentage
      );

      colours.push(myp5_2.color(nameColor.x, nameColor.y, nameColor.z));
    }

    return colours;
  }

  private static getColorByPercentage(
    firstColor: p5.Vector,
    secondColor: p5.Vector,
    percentage: number
  ) {
    // assumes colors are p5js vectors
    var firstColorCopy = firstColor.copy();
    var secondColorCopy = secondColor.copy();

    var deltaColor = secondColorCopy.sub(firstColorCopy);
    var scaledDeltaColor = deltaColor.mult(percentage);
    return firstColorCopy.add(scaledDeltaColor);
  }
}
