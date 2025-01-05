package net.vesperis;

import net.minestom.server.extras.velocity.VelocityProxy;

import java.util.Objects;


public class Universal {

    private final Universal universal;

    public Universal() {
        universal = new Universal();
    }

    public Universal getUniversal() {
        return Objects.requireNonNullElseGet(universal, Universal::new);
    }

    public void enableVelocity(String secret) {
        VelocityProxy.enable(secret);
    }
}