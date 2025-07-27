import 'package:flutter/material.dart';

class HomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Dashboard Provider')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Halo, ini Dashboard Provider!', style: TextStyle(fontSize: 20)),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                // Nanti bisa navigasi ke fitur daftar jasa, dsb.
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Fitur Provider Coming Soon')),
                );
              },
              child: Text('Kelola Jasa Anda'),
            ),
          ],
        ),
      ),
    );
  }
}
