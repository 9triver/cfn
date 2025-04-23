package com.github9triver.cfn.web;

import com.github9triver.cfn.manager.LocalResourceManager;
import com.github9triver.cfn.model.dto.ResourceDto;
import com.github9triver.cfn.model.vo.ResourceVo;
import com.github9triver.cfn.model.vo.Response;
import com.github9triver.cfn.util.mapper.ResourceMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ResourceController {

    private LocalResourceManager localResourceManager;

    @Autowired
    public void setLocalResourceManager(LocalResourceManager localResourceManager) {
        this.localResourceManager = localResourceManager;
    }

    @GetMapping(value = "/resources", produces = MediaType.APPLICATION_JSON_VALUE)
    public Response<ResourceVo> getAllResources() {
        ResourceDto resources = localResourceManager.getAllResources();
        return Response.ok(ResourceMapper.dto2vo(resources));
    }

}
