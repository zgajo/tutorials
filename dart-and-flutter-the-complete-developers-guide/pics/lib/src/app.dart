import 'package:flutter/material.dart';

class App extends StatelessWidget {
  Widget build(context) {
    return MaterialApp(
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
  }
}
