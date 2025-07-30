import 'package:flutter/material.dart';

class HomePage extends StatelessWidget {
   final String username;

  const HomePage({Key? key, required this.username}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Dashboard User')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Halo, ini Dashboard User!', style: TextStyle(fontSize: 20)),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                // Nanti bisa navigasi ke fitur lain
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Fitur User Coming Soon')),
                );
              },
              child: Text('Lihat Jasa Tersedia'),
            ),
          ],
        ),
      ),
    );
  }
}
