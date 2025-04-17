package com.github9triver.cfn;

public interface MessageBuffer {

    void sendTask(Task task);
    Task readTask();

    void sendState(NodeState state);
    NodeState readState();

}
