#include <cstdio>
#include<iostream>
#include <fstream>
#include<string>
using namespace std;
int main(){
	fstream myFile;
	myFile.open("ques1.txt", ios::in);    //for reading file
	if (myFile.is_open()){
		string line;
		while(getline(myFile,line)){
			cout << line << endl;
		}
		myFile.close();
	}

return 0;
}
