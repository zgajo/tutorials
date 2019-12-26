// package: importing file from third party package, not a dart standard lib
// flutter - the name of the package we are importing
// /material.dart - the file we are importing from that package
import 'package:flutter/material.dart';
// imports App widget
import './src/app.dart';

void main() {
  //  creates materialapp widget
  var app = App();

// runApp is imported with material
  runApp(app);
}
