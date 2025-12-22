embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
  position {
    x: -1.0
  }
  scale {
    x: 0.84
    y: 0.84
    z: 0.84
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"enemy\"\n"
  "mask: \"wall\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_CAPSULE\n"
  "    position {\n"
  "      x: -5.0\n"
  "      y: -9.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 2\n"
  "  }\n"
  "  data: 32.227016\n"
  "  data: 17.527594\n"
  "}\n"
  ""
}
