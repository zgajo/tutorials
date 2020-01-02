import 'package:flutter/material.dart';
import 'screens/login.dart';

class App extends StatelessWidget {
  Widget build(context) {
    return MaterialApp(
      title: "Log me in",
      home: Scaffold(
        body: LoginScreen(),
      ),
    );
  }
}
