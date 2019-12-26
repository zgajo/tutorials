// package: importing file from third party package, not a dart standard lib
// flutter - the name of the package we are importing
// /material.dart - the file we are importing from that package
import 'package:flutter/material.dart';

void main() {
  //  creates materialapp widget
  var app = MaterialApp(
    // default route
    home: Scaffold(
      appBar: AppBar(
        title: Text('Pics'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          print('de si ba');
        },
        child: Icon(Icons.add),
      ),
    ),
  );

// runApp is imported with material
  runApp(app);
}
