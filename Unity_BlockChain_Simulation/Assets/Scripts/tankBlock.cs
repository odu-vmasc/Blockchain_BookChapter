using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class tankBlock : MonoBehaviour
{

	public string blockchain;
	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update ()
	{
		if (Input.GetKeyDown(blockchain))
		{
			Renderer rend = GetComponent<Renderer>();
			rend.material.shader = Shader.Find("Standard");
			rend.material.color = Color.blue;
		}
	}
}
