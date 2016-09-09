package com.example;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.SpringApplicationConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.test.context.web.WebAppConfiguration;

@RunWith(SpringJUnit4ClassRunner.class)
@SpringApplicationConfiguration(classes = ZeroDowntimeJreUpdateApplication.class)
@WebAppConfiguration
public class ZeroDowntimeJreUpdateApplicationTests {

	@Test
	public void contextLoads() {
	}

}
