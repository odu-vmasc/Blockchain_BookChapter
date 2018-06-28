using System.Collections;
using System.Collections.Generic;
using System.Net;
using System.Text;  // for class Encoding
using System.IO;    // for StreamReader
using Newtonsoft.Json;
using System.Runtime.InteropServices;
//using NUnit.Framework.Internal.Execution;
using UnityEngine;

public class flight_Circle : MonoBehaviour
{

	public float currentRotation;
	public Vector3 radius;
	public Quaternion rotation;
	public float counter;
	public GameObject minMapColor;
	public string tID;
	public string movekey;
	public float s = 5;
	public GameObject Box;
	
	//private bool flag1;
	
	// Use this for initialization
	void Start ()
	{
		currentRotation = 0;
		radius = new Vector3(5, transform.localPosition.y, transform.localPosition.z);
		rotation = new Quaternion();
		counter = 0.1F;
		
		//flag1 = false;
		InvokeRepeating("verifyBlockChain", 2.0f, 1.0f);

	}
	
	// Update is called once per frame
	void Update ()
	{		
		
		//currentRotation += (transform.localPosition.x+counter) * Time.deltaTime * 100;
		//rotation.eulerAngles = new Vector3(transform.localPosition.x, currentRotation, transform.localPosition.z);
		//transform.position = rotation * radius;
		
		transform.RotateAround(Vector3.zero, Vector3.up, -20 * Time.deltaTime);
		
		
	}
	
	void verifyBlockChain()
	{
		//string _cred = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDk3NzQ1ODEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE1MDk3Mzg1ODF9.c0AbWAJeoUQ5lNjQQQyCe7t-6XaoEG7z43VyP1xCiH0";
		string _cred = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTMzOTA0MTYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE1MTMzNTQ0MTZ9.ExfioRWDVtBOKyUgtAUS5icwZV2SEe2_jQinB6hw1co";
		var request = (HttpWebRequest)WebRequest.Create("http://192.168.121.192:4000/channels/mychannel/chaincodes/mycc1?peer=peer1&fcn=query&args=%5B%22" + tID + "%22%5D");
		//request.Headers[HttpRequestHeader.Authorization] = _cred;
		request.Headers.Add("Authorization", "Bearer " + _cred);
		request.Accept = "application/json";
		var response = (HttpWebResponse)request.GetResponse();			
		var responseString = new StreamReader(response.GetResponseStream()).ReadToEnd();

		string[] splits = responseString.Split(':');
		string mess = splits[3].Split('}')[0];

		if (mess.Contains("Error"))
		{
			Renderer rend = minMapColor.GetComponent<Renderer>();
			rend.material.shader = Shader.Find("Standard");
			rend.material.color = Color.red;// SetColor("_SpecColor", Color.yellow);
		}
		else
		{
			Renderer rend = minMapColor.GetComponent<Renderer>();
			rend.material.shader = Shader.Find("Standard");
			rend.material.color = Color.green;// SetColor("_SpecColor", Color.yellow);
		}
		Debug.Log("!!!!!!!  Blockchain results after the Query!! " + responseString + " : " + splits[3].Split('}')[0]);
	}
}
