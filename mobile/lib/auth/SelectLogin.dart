import 'package:flutter/material.dart';
import 'UserLogin.dart';
import 'ProviderLogin.dart';

class SelectLoginPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xffffd221),
      appBar: AppBar(title: Text('Pilih Login Sebagai')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            ElevatedButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => UserLoginPage()),
                );
              },
              child: Text('Masuk sebagai User'),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => ProviderLoginPage()),
                );
              },
              child: Text('Masuk sebagai Provider'),
            ),
          ],
        ),
      ),
    );
  }
}
