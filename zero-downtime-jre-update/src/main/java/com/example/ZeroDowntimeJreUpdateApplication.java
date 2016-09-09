package com.example;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class ZeroDowntimeJreUpdateApplication {

	@RequestMapping("/")
	String jreVersion() {
		return "System.getProperty(\"java.version\") = " + System.getProperty("java.version");
	}

	public static void main(String[] args) {
		SpringApplication.run(ZeroDowntimeJreUpdateApplication.class, args);
	}
}
