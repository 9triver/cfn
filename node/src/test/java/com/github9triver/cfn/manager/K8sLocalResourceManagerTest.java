package com.github9triver.cfn.manager;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
@SpringBootTest
class K8sLocalResourceManagerTest {

    @Autowired
    private K8sLocalResourceManager k8sLocalResourceManager;

    @Test
    void getAllResources() {
        k8sLocalResourceManager.getAllResources();
    }
}