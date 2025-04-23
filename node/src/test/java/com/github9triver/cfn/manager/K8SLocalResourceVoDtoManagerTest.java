package com.github9triver.cfn.manager;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class K8SLocalResourceVoDtoManagerTest {

    @Autowired
    private LocalResourceManager localResourceManager;

    @Test
    void getAllResources() {
        localResourceManager.getAllResources();
    }
}