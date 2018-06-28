using System.Collections;
using System.Collections.Generic;
using System.Net;
using Newtonsoft.Json;
using System.IO;
using System.Net.Sockets;
using UnityEngine;

public class byzantineScript : MonoBehaviour
{

	public string activateByz;
	
	// Use this for initialization
	void Start () {
		InvokeRepeating("verifyBlockChain", 2.0f, 1.0f);	
	}

	void verifyBlockChain()
	{
		//string _cred = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDk3NzQ1ODEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE1MDk3Mzg1ODF9.c0AbWAJeoUQ5lNjQQQyCe7t-6XaoEG7z43VyP1xCiH0";
		string _cred = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTMyNDk4NzcsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE1MTMyMTM4Nzd9.SOr5ffSJbHqVW2ZULq1NATk-GEeKa8KMZ1HddafJaNc";
		//var request = (HttpWebRequest)WebRequest.Create("http://192.168.121.192:4000/channels/mychannel/chaincodes/mycc7?peer=peer1&fcn=query&args=%5B%22" + tID + "%22%5D");
		//var request = (HttpWebRequest) WebRequest.Create("http://192.168.121.192:7063");
		//request.Headers[HttpRequestHeader.Authorization] = _cred;
		//request.Headers.Add("Authorization", "Bearer " + _cred);
		//request.Accept = "application/json";
		//var response = (HttpWebResponse)request.GetResponse();			
		//var responseString = new StreamReader(response.GetResponseStream()).ReadToEnd();

		using (TcpClient tcpClient = new TcpClient())
		{
			try
			{
				tcpClient.Connect("192.168.121.192", 7063);
				Renderer rend = GetComponent<Renderer>();
				rend.material.shader = Shader.Find("Standard");
				rend.material.color = Color.green;// SetColor("_SpecColor", Color.yellow);
				Debug.Log("!!!!!!!  Peer port is open and is active!! ");
			}
			catch
			{
				Renderer rend = GetComponent<Renderer>();
				rend.material.shader = Shader.Find("Standard");
				rend.material.color = Color.yellow;// SetColor("_SpecColor", Color.yellow);
				Debug.Log("!!!!!!!  Peer port is closed!! ");
			}
		}
		//string[] splits = responseString.Split(':');
		//string mess = splits[3].Split('}')[0];

		/*
		if (responseString.Contains("Failed"))
		{

		}
		else
		{

		}
		*/
		
	}
	
	// Update is called once per frame
	void Update () {

		if (Input.GetKeyDown(activateByz))
		{
			Renderer rend = GetComponent<Renderer>();
			rend.material.shader = Shader.Find("Standard");
			rend.material.color = Color.yellow;// SetColor("_SpecColor", Color.yellow);
			
		}
	}
}
