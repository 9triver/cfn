package com.github9triver.cfn.web;

import com.github9triver.cfn.manager.LocalResourceManager;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ResourceController {

    private LocalResourceManager localResourceManager;

    @Autowired
    public void setLocalResourceManager(LocalResourceManager localResourceManager) {
        this.localResourceManager = localResourceManager;
    }

    @GetMapping("/resources")
    public Object getAllResources() {
        return localResourceManager.getAllResources();
    }

}
