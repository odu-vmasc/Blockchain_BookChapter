using System.Collections;
using System.Collections.Generic;

using UnityEngine;

public class ai_move : MonoBehaviour
{

	public float speed = 5;
	public float directionChangeInterval = 1;
	public float maxHeadingChange = 30;
	public float rotateSpeed = 3.0F;
	public GameObject byzBox;
	
	CharacterController controller;
	float heading;
	private Vector3 targetRotation;
	private bool flag1, flag2;
	
	
	// Use this for initialization
	void Start ()
	{
		flag1 = false;
		flag2 = false;
	}
	
	// Update is called once per frame
	void Update ()
	{		
		//transform.eulerAngles =
		//	Vector3.Slerp(transform.eulerAngles, targetRotation, Time.deltaTime * directionChangeInterval);
		//controller.SimpleMove(Vector3.forward * speed);
		controller = GetComponent<CharacterController>();
		//transform.Rotate(0, Input.GetAxis("Horizontal") * rotateSpeed, 0);
		Vector3 direction = new Vector3(0, 0, 1);
		Vector3 forward = transform.TransformDirection(direction);
		float curSpeed = speed;// * Input.GetAxis("Vertical");
		//controller.SimpleMove(forward * curSpeed);

		if (Input.GetKeyDown("space"))
		{
			if (flag1)
			{
				flag2 = true;
			}
			else
			{
				flag1 = true;
			}
		}
		
		if (flag1 && transform.localPosition.z > byzBox.transform.position.z)
		{
			transform.localPosition += transform.forward * curSpeed;
		}
	}
}
