components {
  id: "enemy"
  component: "/main/enemy/enemy.script"
  position {
    x: 88.14124
    y: 36.89382
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    x: 30.0
    y: 14.0
  }
  scale {
    x: 0.131023
    y: 0.131023
    z: 0.131023
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 29.0\n"
  "      y: 14.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 26.909166\n"
  "  data: 14.130273\n"
  "  data: 10.533334\n"
  "}\n"
  ""
}
